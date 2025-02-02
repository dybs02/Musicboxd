import { ApolloClient, ApolloLink, InMemoryCache, concat, createHttpLink } from '@apollo/client/core';
import { useAuthStore } from './authStore';


const httpLink = createHttpLink({
  // @ts-ignore
  uri: import.meta.env.VITE_GRAPHQL_API_URL,
  credentials: 'include',
});


const authMiddleware = new ApolloLink((operation, forward) => {
  const store = useAuthStore();
  const token = store.getJWT();
  operation.setContext({
    headers: {
      Authorization: token ? `${token}` : "",
    },
  });
  return forward(operation);
});


const apolloClient = new ApolloClient({
  link: concat(authMiddleware, httpLink),
  cache: new InMemoryCache(),
  credentials: 'include',
});


export default apolloClient;