import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import UserAuthenticationPage from './pages/userAuthentication/userAuthenticationPage';
import { createTheme, ThemeProvider } from '@mui/material';
import { Provider } from 'react-redux';
import store from './redux/store';
import LandingPage from './pages/landing/landing';

const router = createBrowserRouter([
  {
    path: "/",
    element: <LandingPage />
  },
  {
    path: "/authenticate",
    element: <UserAuthenticationPage />
  },
]);

const theme = createTheme({
  palette: {
    primary: {
      main: "#49BBBD",
    },
    secondary: {
      main: 'rgba(73, 187, 189, 0.4)',
    },
    info: {
      main: '#FFFFFF'
    }
  }
});


ReactDOM.createRoot(document.getElementById('root')!).render(
  <Provider store={store}>
    <ThemeProvider theme={theme}>
      <React.StrictMode>
        <RouterProvider router={router} />
      </React.StrictMode>
    </ThemeProvider>
  </Provider>
)
