import { emptyUser, type UserType } from "./user"



type ChatType = {
  _id: string
  name: string
  participantsIds: string[]
  participants: UserType[]
  messages: MessageType[]
  createdAt: string
}

const emptyChat: ChatType = {
  _id: '',
  name: '',
  participantsIds: [],
  participants: [],
  messages: [],
  createdAt: ''
}

type MessageType = {
  _id: string
  content: string
  senderId: string
  sender: UserType
  createdAt: string
}

const emptyMessage: MessageType = {
  _id: '',
  content: '',
  senderId: '',
  sender: emptyUser,
  createdAt: ''
}

type MessagesPage = {
  totalMessages: number
  totalPages: number
  hasPreviousPage: boolean
  hasNextPage: boolean
  messages: MessageType[]
}

const emptyMessagesPage: MessagesPage = {
  totalMessages: 0,
  totalPages: 0,
  hasPreviousPage: false,
  hasNextPage: false,
  messages: []
}

export type {
  ChatType,
  MessageType,
  MessagesPage,
};

export {
  emptyChat,
  emptyMessage,
  emptyMessagesPage,
};