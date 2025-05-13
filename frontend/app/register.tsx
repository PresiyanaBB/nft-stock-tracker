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

export default function Register() {
  const { authenticate, isLoadingAuth } = useAuth();
  const navigation = useNavigation();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');

  async function onRegister() {
    if (!email.trim() || !password.trim() || !confirmPassword.trim()) {
      alert("All fields are required.");
      return;
    }

    if (password !== confirmPassword) {
      alert("Passwords do not match.");
      return;
    }

    console.log("Register button pressed");
    try {
      await authenticate("register", email, password);
      router.push('/(nft-purchase)');
    } catch (error) {
      console.error("Registration Error:", error);
      alert("Registration failed. Please try again.");
    }
  }

  function onGoToLogin() {
    console.log("Redirect to login");
    router.push('/login');
  }

  return (
    <KeyboardAvoidingView behavior="padding" style={globals.container}>
      <ScrollView contentContainerStyle={globals.container}>
        <VStack flex={1} justifyContent='center' alignItems='center' p={40} gap={40}>
          <HStack gap={10}>
            <Text fontSize={30} bold mb={20}>NFT Shop</Text>
            <TabBarIcon name="person-circle-outline" size={50} />
          </HStack>

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

            <VStack gap={5}>
              <Text ml={10} fontSize={14} color="gray">Confirm Password</Text>
              <Input
                value={confirmPassword}
                onChangeText={setConfirmPassword}
                secureTextEntry
                placeholder="Confirm Password"
                placeholderTextColor="darkgray"
                autoCapitalize="none"
                autoCorrect={false}
                h={48}
                p={14}
              />
            </VStack>

            <Button isLoading={isLoadingAuth} onPress={onRegister}>Register</Button>
          </VStack>

          <Divider w={"90%"} />

          <Text onPress={onGoToLogin} fontSize={16} underline>
            Already have an account? Login here
          </Text>
        </VStack>
      </ScrollView>
    </KeyboardAvoidingView>
  );
}
