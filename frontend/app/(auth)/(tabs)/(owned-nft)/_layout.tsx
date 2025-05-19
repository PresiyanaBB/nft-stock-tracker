import { Stack } from 'expo-router';

export default function UserNFTsLayout() {
  return (
    <Stack screenOptions={ { headerBackTitle: "UserNFTs" } }>
      <Stack.Screen name="index" />
      <Stack.Screen name="nft/[id]" />
    </Stack>
  );
}