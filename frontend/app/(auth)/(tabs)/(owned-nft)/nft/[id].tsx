import { Text } from "@/components/Text";
import { VStack } from "@/components/VStack";
import { userNFTService } from "@/services/userNFTs";
import { NFT } from "@/types/NFT";
import { useFocusEffect } from "@react-navigation/native";
import { router, useLocalSearchParams, useNavigation } from "expo-router";
import { useCallback, useEffect, useState } from "react";
import { Alert, Image } from "react-native";
import { validate as validateUUID } from "uuid";

export default function UserNFTDetailScreen() {
  const navigation = useNavigation();
  const { id } = useLocalSearchParams();
  const [NFT, setNFT] = useState<NFT | null>(null);
  const [qrcode, setQrcode] = useState<string | null>(null);

  async function fetchUserNFT() {
    if (!id || !validateUUID(id.toString())) {
      Alert.alert("Error", "Invalid NFT ID.");
      router.back();
      return;
    }

    try {
      const userNFTs = await userNFTService.getUserNFTs();
      const userNFTId = userNFTs.data.find((userNFT) => {
        return userNFT.nft_id === id;
      });
      const { data } = await userNFTService.getUserNFT(userNFTId.id);
      setNFT(data.userNFT.nft);
      setQrcode(data.qrcode);
    } catch (error) {
      console.error("Error fetching user NFT:", error);
      Alert.alert("Error", "Failed to load NFT details.");
      router.back();
    }
  }

  useFocusEffect(
    useCallback(() => {
      fetchUserNFT();
    }, [])
  );

  useEffect(() => {
    navigation.setOptions({ headerTitle: NFT ? NFT.name : "NFT Details" });
  }, [navigation, NFT]);

  if (!NFT) return null;

  return (
    <VStack
      alignItems="center"
      m={20}
      p={20}
      gap={20}
      flex={1}
      style={{
        backgroundColor: "white",
        borderRadius: 20,
        dowColor: "#000",
        borderRadius: 20,
        boxShadow: "0px 4px 10px rgba(0, 0, 0, 0.2)",
        elevation: 5,
      }}
    >
      <Text fontSize={22} bold>{NFT.name}</Text>
      {qrcode ? (
        <Image
          style={{ width: 300, height: 300, borderRadius: 10 }}
          source={{ uri: `data:image/png;base64,${qrcode}` }}
        />
      ) : (
        <Text fontSize={16} color="gray">QR Code not available</Text>
      )}
    </VStack>
  );
}
