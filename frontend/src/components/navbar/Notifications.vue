<script setup lang="ts">
import { useAuthStore } from "@/services/authStore";
import { MARK_NOTIFICATION_AS_READ, NOTIFICATIONS_PAGE } from "@/services/queries";
import type { NotificationType } from '@/types/notifications';
import { handleGqlError } from "@/utils/error";
import { navigateToChat } from "@/utils/navigate";
import { useMutation, useQuery, useSubscription } from '@vue/apollo-composable';
import Badge from 'primevue/badge';
import Button from 'primevue/button';
import OverlayPanel from 'primevue/overlaypanel';
import { ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';
import { useRoute, useRouter } from 'vue-router';



const { locale, t } = useI18n()
const store = useAuthStore();
const route = useRoute();
const router = useRouter();

const notifications = ref<NotificationType[]>([]);
const unreadCount = ref(0);
const op = ref();


const fetchNotifications = () => {
  const { loading, onError, onResult } = useQuery(
    NOTIFICATIONS_PAGE,
    {
      pageSize: 10,
      page: 1
    }
  );

  onError((err) => {
    console.error('Error fetching notifications:', err);
    // handleGqlError(router, err);
  });
  
  onResult((res: any) => {
    if (res.loading) {
      return;
    }

    console.log('Notifications result:', res);
    if (res.data && res.data.notificationsPage) {
      // TODO use pagination
      notifications.value = res.data.notificationsPage.notifications;
      unreadCount.value = notifications.value.filter((n: NotificationType) => !n.isRead).length;
    }
  })
};

const toggleNotifications = (event: any) => {
  op.value.toggle(event);
};

const markAllAsRead = () => {
  // TODO: Implement mutation to mark all notifications as read
  unreadCount.value = 0;
  notifications.value.forEach((n: NotificationType) => n.isRead = true);
};





const { mutate: markAsRead, error: markAsReadError, onDone: markAsReadOnDone } = useMutation(
  MARK_NOTIFICATION_AS_READ
);

watch(markAsReadError, (err) => {
  handleGqlError(router, err, ["mongo: no documents in result"]);
});


const openNotification = async (notification: NotificationType) => {
  if (!notification.isRead) {
    await markAsRead({
      notificationId: notification._id,
    });
  }
    
  navigateToChat(router, JSON.parse(notification.message).chatWith);
  op.value.hide();
};


watch(() => route.params, fetchNotifications, { immediate: true })

</script>

<template>
  <Button
    @click="toggleNotifications"
    class="p-button-rounded p-button-text"
    icon="pi pi-bell"
    >
    <Badge v-if="unreadCount > 0" :value="unreadCount"></Badge>
  </Button>
  <OverlayPanel ref="op" appendTo="body" :showCloseIcon="true" id="overlay_panel" style="width: 450px">
    <div class="flex justify-content-between align-items-center mb-3">
      <span class="text-xl font-bold">Notifications</span>
      <!-- <Button icon="pi pi-check" label="Mark all as read" class="p-button-text p-button-sm" @click="markAllAsRead" /> -->
    </div>
  <ul class="list-none p-0 m-0 space-y-1">
    <li v-for="notification in notifications" :key="notification._id" class="flex items-center p-3 rounded-lg border hover:bg-opacity-50 transition-colors cursor-pointer" @click="openNotification(notification)">
      <i :class="['pi pi-envelope mr-3 text-lg', { 'opacity-100': !notification.isRead, 'opacity-40': notification.isRead }]"></i>
      <div class="flex flex-col">
        <span :class="{ 'font-semibold': !notification.isRead, 'opacity-70': notification.isRead }">{{ t('newChatMessageInChat') + ' ' + JSON.parse(notification.message).chatName }}</span>
        <span class="text-sm opacity-60">{{ new Date(notification.createdAt).toLocaleString() }}</span>
      </div>
    </li>
    <li v-if="notifications.length === 0" class="p-4 text-center opacity-60">{{ t('noNotifications') }}</li>
  </ul>
  </OverlayPanel>
</template>

<style scoped>

</style>
