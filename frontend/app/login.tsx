import { useState } from 'react';
import { Button } from '@/components/Button';
import { Divider } from '@/components/Divider';
import { HStack } from '@/components/HStack';
import { Input } from '@/components/Input';
import { Text } from '@/components/Text';
import { VStack } from '@/components/VStack';
import { TabBarIcon } from '@/components/navigation/TabBarIcon';
import { useAuth } from '@/context/AuthContext';
import { KeyboardAvoidingView, ScrollView } from 'react-native';
import { globals } from '@/styles/_global';
import { useNavigation, router } from 'expo-router';

export default function Login() {
  const { authenticate, isLoadingAuth } = useAuth();
  const navigation = useNavigation();

  const [authMode, setAuthMode] = useState<"login" | "register">('login');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  async function onAuthenticate() {
    if (!email.trim() || !password.trim()) {
      alert("Email and Password are required.");
      return;
    }

    console.log("Login button pressed");
    try {
      await authenticate(authMode, email, password);
      if (authMode === "login") {
        router.push('/(owned-nft)');
      } else {
        router.push('/register');
      }
    } catch (error) {
      console.error("Authentication Error:", error);
      alert("Authentication failed. Please try again.");
    }
  }

  function onToggleAuthMode() {
    router.push('/register');
    setAuthMode(authMode === 'login' ? 'register' : 'login');
  }

  return (
    <KeyboardAvoidingView behavior="padding" style={globals.container}>
      <ScrollView contentContainerStyle={globals.container}>
        <VStack flex={1} justifyContent='center' alignItems='center' p={40} gap={40}>

          <HStack gap={10}>
            <Text fontSize={30} bold mb={20}>NFT shop</Text>
            <TabBarIcon name="person-circle-outline" size={50} />
          </HStack >

          <VStack w={"100%"} gap={30}>

            <VStack gap={5}>
              <Text ml={10} fontSize={14} color="gray">Email</Text>
              <Input
                value={email}
                onChangeText={setEmail}
                placeholder="Email"
                placeholderTextColor="darkgray"
                autoCapitalize="none"
                autoCorrect={false}
                h={48}
                p={14}
              />
            </VStack>

            <VStack gap={5}>
              <Text ml={10} fontSize={14} color="gray">Password</Text>
              <Input
                value={password}
                onChangeText={setPassword}
                secureTextEntry
                placeholder="Password"
                placeholderTextColor="darkgray"
                autoCapitalize="none"
                autoCorrect={false}
                h={48}
                p={14}
              />
            </VStack>

            <Button isLoading={isLoadingAuth} onPress={onAuthenticate}>{authMode}</Button>

          </VStack>

          <Divider w={"90%"} />

          <Text onPress={onToggleAuthMode} fontSize={16} underline>
            {authMode === 'login' ? 'Register new account' : 'Login to account'}
          </Text>
        </VStack>
      </ScrollView>
    </KeyboardAvoidingView>
  );
}
