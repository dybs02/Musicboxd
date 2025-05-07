import type { ApolloError } from '@apollo/client/core/index.js';
import type { Router } from 'vue-router';



const handleGqlError = (router: Router, error: ApolloError | null, exceptionErrorMessages: string[] = []) => {
  if (error === null) {
    return;
  }

  if (exceptionErrorMessages.includes(error.message)) {
    return;
  }


  switch (error.message) {
    case 'GraphQL error: Not authenticated':
      // TODO navigate to login page
      console.error('Not authenticated');
      break;
    case 'GraphQL error: Not authorized':
      // TODO navigate to error page
      console.error('Not authorized');
      break;
    case 'GraphQL error: User not found':
      // TODO navigate to error page
      console.error('User not found');
  }


  // TODO disable in PROD
  router.push({ 
    name: 'error', 
    params: {
      message: error.message,
      cause: error.cause ? JSON.stringify(error.cause) : '',
    },
  });
}


export {
  handleGqlError,
}