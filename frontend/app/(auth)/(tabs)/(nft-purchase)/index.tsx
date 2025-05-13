import { HStack } from "@/components/HStack";
import { Text } from "@/components/Text";
import { VStack } from "@/components/VStack";
import { userNFTService } from "@/services/userNFTs";
import { UserNFT } from "@/types/userNFT";
import { useFocusEffect } from "@react-navigation/native";
import { router, useNavigation } from "expo-router";
import { useCallback, useEffect, useState } from "react";
import { Alert, FlatList, TouchableOpacity, Image, Button } from "react-native";
import { validate as validateUUID } from 'uuid';
import { useAuth } from "@/context/AuthContext";
import { UserRole } from "@/types/user";

export default function UserNFTScreen() {
  const { user } = useAuth();
  const navigation = useNavigation();
  const [isLoading, setIsLoading] = useState(false);
  const [NFTs, setNFTs] = useState<UserNFT[]>([]);

  function onGoToUserNFTPage(id: string) {
    if (validateUUID(id)) {
      router.push(`/(nft-purchase)/nft/${id}`);
    } else {
      Alert.alert("Error", "Invalid NFT ID.");
    }
  }

  async function fetchNFTs() {
    try {
      setIsLoading(true);
      const response = await userNFTService.getAll();
      setNFTs(response.data);
    } catch (error) {
      console.error("Error fetching NFTs:", error);
      Alert.alert("Error", "Failed to fetch NFTs");
    } finally {
      setIsLoading(false);
    }
  }

  useFocusEffect(
    useCallback(() => {
      fetchNFTs();
    }, [])
  );

  useEffect(() => {
    navigation.setOptions({ headerTitle: "NFTs" });
  }, [navigation]);

  return (
    <VStack flex={1} p={20} gap={20}>
      <HStack alignItems="center" justifyContent="space-between">
        <Text fontSize={18} bold>{NFTs.length} NFTs</Text>
      </HStack>

      {/* Create NFT Section for Admins */}
      {user?.role === UserRole.Admin && (
        <VStack mb={20}>
          <TouchableOpacity
            style={{
              backgroundColor: "dodgerblue",
              padding: 15,
              borderRadius: 10,
              alignItems: "center",
            }}
            onPress={() => router.push('/(nft-creation)')}
          >
            <Text style={{ color: "white", fontSize: 16, fontWeight: "bold" }}>
              Create NFT
            </Text>
          </TouchableOpacity>
        </VStack>
      )}

      <FlatList
        keyExtractor={({ id }) => id.toString()}
        data={NFTs}
        onRefresh={fetchNFTs}
        refreshing={isLoading}
        ItemSeparatorComponent={() => <VStack h={20} />}
        renderItem={({ item: NFT }) => (
          <TouchableOpacity
            disabled={NFT.collected}
            onPress={() => onGoToUserNFTPage(NFT.id)}
          >
          <VStack
             gap={12}
             h={120}
             style={{
               opacity: NFT.collected ? 0.5 : 1,
               backgroundColor: "white",
               borderRadius: 20,
               boxShadow: "0px 2px 5px rgba(0, 0, 0, 0.1)",
               padding: 15,
             }}
           >

              <HStack alignItems="center" justifyContent="space-between">
                <HStack alignItems="center" gap={10}>
                  <Image
                    source={{ uri: `data:image/png;base64,${NFT.NFT.image}` }}
                    style={{ width: 50, height: 50, borderRadius: 10 }}
                    resizeMode="cover"
                  />
                  <VStack>
                    <Text fontSize={18} bold>{NFT.NFT.name}</Text>
                    <Text fontSize={12} color="gray">
                      {new Date(NFT.NFT.createdAt).toLocaleDateString()}
                    </Text>
                  </VStack>
                </HStack>

                <VStack alignItems="center">
                  <Text fontSize={16} bold>
                    {NFT.collected ? "Owned" : "Not-Owned"}
                  </Text>
                  {NFT.collected && (
                    <Text mt={4} fontSize={10} color="gray">
                      {new Date(NFT.updatedAt).toLocaleString()}
                    </Text>
                  )}
                </VStack>
              </HStack>
            </VStack>
          </TouchableOpacity>
        )}
      />
    </VStack>
  );
}
