import { HStack } from "@/components/HStack";
import { Text } from "@/components/Text";
import { VStack } from "@/components/VStack";
import { userNFTService } from "@/services/userNFTs";
import { NFT } from "@/types/NFT";
import { NFTService } from "@/services/NFTs";
import { useFocusEffect } from "@react-navigation/native";
import { router, useNavigation } from "expo-router";
import { useCallback, useEffect, useState } from "react";
import { Image, Alert, FlatList, TouchableOpacity } from "react-native";
import { validate as validateUUID } from 'uuid';
import { useAuth } from "@/context/AuthContext";
import { UserRole } from "@/types/user";

export default function UserNFTScreen() {
  const { user } = useAuth();
  const navigation = useNavigation();
  const [isLoading, setIsLoading] = useState(false);
  const [NFTs, setNFTs] = useState<NFT[]>([]);

  function onGoToUserNFTPage(id: string) {
    if (validateUUID(id)) {
      router.push(`/(owned-nft)/nft/${id}`);
    } else {
      Alert.alert("Error", "Invalid NFT ID.");
    }
  }

  async function fetchNFTs() {
    try {
      setIsLoading(true);
      const response = await NFTService.getAll();
      const userNFTs = await userNFTService.getUserNFTs();
      const allUserNFTs = await userNFTService.getAll();
      response.data = response.data.filter((nft: NFT) => {
        nft.collected = false;
        return userNFTs.data.some((userNFT) => {
          if (userNFT.nft_id === nft.id) {
            nft.collected = userNFT.collected;
          }
          return userNFT.nft_id === nft.id;
        })
      });
      console.log("All User NFTs", allUserNFTs.data);
      console.log("User NFTs", userNFTs.data);
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
    navigation.setOptions({
      headerTitle: "NFTs",
      headerRight: () => < Text >Welcome, {user?.email}   </Text >
    });
  }, [navigation]);

  return (
    <VStack flex={1} p={20} gap={20}>
      <HStack alignItems="center" justifyContent="space-between">
        <Text fontSize={18} bold> {NFTs.length} NFTs</Text>
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
            onPress={() => router.push('/(buy-nft)/new')}
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
        renderItem={({ item: current_NFT }) => (
          <TouchableOpacity
            disabled={current_NFT.collected}
            onPress={() => onGoToUserNFTPage(current_NFT.id)}
          >
            <VStack
              gap={12}
              h={120}
              style={{
                opacity: current_NFT.collected ? 0.5 : 1,
                backgroundColor: "white",
                borderRadius: 20,
                boxShadow: "0px 2px 5px rgba(0, 0, 0, 0.1)",
                padding: 15,
              }}
            >
              <HStack alignItems="center" justifyContent="space-between">
                <HStack alignItems="center" gap={10}>
                  <VStack>
                    <Image source={{ uri: `data:image/png;base64,${current_NFT.image}` }}
                      style={{ width: 100, height: 100, borderRadius: 10 }}
                      resizeMode="cover"
                    />
                  </VStack>
                  <VStack>
                    <Text fontSize={18} bold>{current_NFT.name}</Text>
                    <Text fontSize={22} bold> | </Text>
                    <Text fontSize={18} bold>{current_NFT.price} $</Text>
                  </VStack>
                </HStack>

                <VStack alignItems="center">
                  <Text fontSize={16} bold>
                    {current_NFT.collected ? "Owned" : "Not-Owned"}
                  </Text>
                  {current_NFT.collected && (
                    <Text mt={4} fontSize={10} color="gray">
                      Created on: {current_NFT.created_at.split("T")[0]}
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
