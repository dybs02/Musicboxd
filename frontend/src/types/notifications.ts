type NotificationType = {
  _id: string;
  message: string;
  type: string;
  notifyingUserId: string;
  notifiedUserId: string;
  isRead: boolean;
  createdAt: string;
};

const emptyNotification: NotificationType = {
  _id: "",
  message: "",
  type: "",
  notifyingUserId: "",
  notifiedUserId: "",
  isRead: false,
  createdAt: "",
};

type NotificationsPageType = {
  totalNotifications: number;
  totalPages: number;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
  notifications: NotificationType[];
};

const emptyNotificationsPage: NotificationsPageType = {
  totalNotifications: 0,
  totalPages: 0,
  hasPreviousPage: false,
  hasNextPage: false,
  notifications: [],
};

export type {
  NotificationType,
  NotificationsPageType,
};

export {
  emptyNotification,
  emptyNotificationsPage,
};
