import { emptyUser, type UserType } from "./user"



type ChatType = {
  _id: string
  name: string
  participantsIds: string[]
  participants: UserType[]
  participantId: string
  participant: UserType
  messages: MessageType[]
  createdAt: string
}

const emptyChat: ChatType = {
  _id: '',
  name: '',
  participantsIds: [],
  participants: [],
  participantId: '',
  participant: emptyUser,
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

export type {
  ChatType,
  MessageType,
};

export {
  emptyChat,
  emptyMessage,
};