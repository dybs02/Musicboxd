type Comments = {
  text: string;
  createdAt: string;
  updatedAt: string;
  user: {
    _id: string;
    displayName: string;
    images: {
      url: string;
    }[];
  };
}[];
