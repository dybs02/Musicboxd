import { ApolloClient, ApolloLink, InMemoryCache, concat, createHttpLink } from '@apollo/client/core';


const httpLink = createHttpLink({
  // @ts-ignore
  uri: import.meta.env.VITE_GRAPHQL_API_URL,
  credentials: 'include',
});


const authMiddleware = new ApolloLink((operation, forward) => {
  // TODO set token
  const token = null;
  operation.setContext({
    headers: {
      authorization: token ? `Bearer ${token}` : "",
    },
  });
  return forward(operation);
});


const apolloClient = new ApolloClient({
  link: concat(authMiddleware, httpLink),
  cache: new InMemoryCache(),
});


export default apolloClient;