import { Stack } from 'expo-router';

export default function NFTsLayout() {
  return (
    <Stack screenOptions={ { headerBackTitle: "NFTs" } }>
      <Stack.Screen name="index" />
      <Stack.Screen name="new" />
      <Stack.Screen name="nft/[id]" />
    </Stack>
  );
}