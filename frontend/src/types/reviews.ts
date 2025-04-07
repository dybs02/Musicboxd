export type Review = {
  _id: number;
  value: number;
  itemId: string;
  title: string;
  description: string;
  userId: string;
  album: {
    name: string;
    images: {
      url: string;
      width: number;
      height: number;
    }[];
    artists: {
      external_urls: {
        spotify: string;
      };
      href: string;
      id: string;
      name: string;
      type: string;
      uri: string;
    }[];
  };
};

const emptyReview: Review = {
  _id: 0,
  value: 0,
  itemId: "",
  title: "",
  description: "",
  userId: "",
  album: {
    name: "",
    images: [
      {
        url: "",
        width: 0,
        height: 0
      }
    ],
    artists: [
      {
        external_urls: {
          spotify: ""
        },
        href: "",
        id: "",
        name: "",
        type: "",
        uri: ""
      }
    ]
  }
};

export { emptyReview };