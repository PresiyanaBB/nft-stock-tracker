// AuthContext.tsx
import { userService } from '@/services/user';
import AsyncStorage from '@react-native-async-storage/async-storage';
import React, { useEffect, useState } from 'react';
import { router } from 'expo-router';
import { User } from '@/types/user';
import axios from 'axios';

interface AuthContextProps {
  isLoggedIn: boolean;
  isLoadingAuth: boolean;
  authenticate: (authMode: "login" | "register", email: string, password: string) => Promise<void>;
  logout: VoidFunction;
  user: User | null;
}

const AuthContext = React.createContext({} as AuthContextProps);

export function useAuth() {
  return React.useContext(AuthContext);
}

export function AuthenticationProvider({ children }: React.PropsWithChildren) {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [isLoadingAuth, setIsLoadingAuth] = useState(false);
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    async function checkIfLoggedIn() {
      const token = await AsyncStorage.getItem('token');
      const user = await AsyncStorage.getItem('user');

      if (token && user) {
        setIsLoggedIn(true);
        setUser(JSON.parse(user));
        router.replace('(auth)');
      } else {
        setIsLoggedIn(false);
      }
    }

    checkIfLoggedIn();
  }, []);

  async function authenticate(authMode: "login" | "register", email: string, password: string) {
    try {
      setIsLoadingAuth(true);

      const response = await userService[authMode](email, password);

      if (response) {
        setIsLoggedIn(true);
        await AsyncStorage.setItem('token', response.token);
        await AsyncStorage.setItem('user', JSON.stringify(response.user));
        setUser(response.user);
        router.replace('(auth)');
      }
    } catch (error) {
      setIsLoggedIn(false);

      if (axios.isAxiosError(error)) {
        console.error("Authentication Error:", error.response?.data);
        alert(error.response?.data?.error || "Authentication failed. Please try again.");
      } else {
        console.error("Unexpected Error:", error);
        alert("An unexpected error occurred. Please try again.");
      }
    } finally {
      setIsLoadingAuth(false);
    }
  }

  async function logout() {
    setIsLoggedIn(false);
    await AsyncStorage.removeItem('token');
    await AsyncStorage.removeItem('user');
    setUser(null);
    router.replace('/login');
  }

  return (
    <AuthContext.Provider
      value={{
        authenticate,
        logout,
        isLoggedIn,
        isLoadingAuth,
        user,
      }}>
      {children}
    </AuthContext.Provider>
  );
}
