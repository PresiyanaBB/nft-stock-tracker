import { Button } from '@/components/Button';
import { Divider } from '@/components/Divider';
import { HStack } from '@/components/HStack';
import { Text } from '@/components/Text';
import { VStack } from '@/components/VStack';
import { TabBarIcon } from '@/components/navigation/TabBarIcon';
import { useAuth } from '@/context/AuthContext';
import { NFTService } from '@/services/NFTs';
import { userNFTService } from '@/services/userNFTs';
import { NFT } from '@/types/nft';
import { UserRole } from '@/types/user';
import { useFocusEffect } from '@react-navigation/native';
import { useNavigation, router } from 'expo-router';
import { useCallback, useEffect, useState } from 'react';
import { Alert, FlatList, TouchableOpacity } from 'react-native';
import { validate as validateUUID } from 'uuid';
import { Image } from 'react-native';


export default function NFTsScreen() {
  const { user } = useAuth();
  const navigation = useNavigation();
  const [isLoading, setIsLoading] = useState(false);
  const [NFTs, setNFTs] = useState<NFT[]>([]);

  useEffect(() => {
    navigation.setOptions({
      headerTitle: "NFTs",
      headerRight: () => (user?.role === UserRole.Admin ? <HeaderRight /> : null),
    });
  }, [navigation, user]);


  const HeaderRight = () => {
  const { user } = useAuth();

  function goToNewNFT() {
    if (user?.role === UserRole.Admin) {
      router.push('/(nft-creation)/new');
    } else {
      Alert.alert("Access Denied", "Only administrators can create NFTs.");
    }
  }

  return (
    <TabBarIcon
      size={32}
      name="add-circle-outline"
      onPress={goToNewNFT}
    />
    );
  };


  function onGoToNFTPage(id: string) {
    if (user?.role === UserRole.Admin) {
      router.push(`/(nft-creation)/nft/${id}`);
    }
  }

  async function buyNFT(id: string) {
    if (!validateUUID(id)) {
      Alert.alert("Error", "Invalid NFT ID");
      return;
    }

    try {
      await userNFTService.createNFT(id);
      Alert.alert("Success", "NFT purchased successfully");
      fetchNFTs();
    } catch (error) {
      Alert.alert("Error", "Failed to buy NFT");
    }
  }

  const fetchNFTs = async () => {
    try {
      setIsLoading(true);
      const response = await NFTService.getAll();
      setNFTs(response.data);
    } catch (error) {
      Alert.alert("Error", "Failed to fetch NFTs");
    } finally {
      setIsLoading(false);
    }
  };

  useFocusEffect(
    useCallback(() => {
      fetchNFTs();
    }, [])
  );

  useEffect(() => {
    navigation.setOptions({
      headerTitle: "NFTs",
      headerRight: user?.role === UserRole.Admin ? HeaderRight : null,
    });
  }, [navigation, user]);

  return (
    <VStack flex={1} p={20} pb={0} gap={20}>
      <HStack alignItems='center' justifyContent='space-between'>
        <Text fontSize={18} bold>{NFTs.length} NFTs</Text>
      </HStack>

      <FlatList
        keyExtractor={(item) => item.id.toString()}
        data={NFTs}
        onRefresh={fetchNFTs}
        refreshing={isLoading}
        ItemSeparatorComponent={() => <VStack h={20} />}
        renderItem={({ item: NFT }) => (
          <VStack
            gap={20}
            p={20}
            style={{
              backgroundColor: 'white',
              borderRadius: 20,
            }}
          >
            <TouchableOpacity onPress={() => onGoToNFTPage(NFT.id)}>
              <HStack alignItems='center' justifyContent="space-between">
                <HStack alignItems='center'>
                  <Text fontSize={26} bold>{NFT.name}</Text>
                  <Text fontSize={26} bold> | </Text>
                  <Text fontSize={16} bold>{NFT.price}</Text>
                </HStack>
                {user?.role === UserRole.Admin && (
                  <TabBarIcon size={24} name="chevron-forward" />
                )}
              </HStack>
            </TouchableOpacity>

            <Divider />

            <Image source={{ uri: `data:image/png;base64,${NFT.image}` }}
                   style={{ width: 100, height: 100, borderRadius: 10 }}
                   resizeMode="cover"
            />

            <Divider />

            <HStack justifyContent='space-between'>
              <Text bold fontSize={16} color='gray'>
                Sold: {NFT.totalNFTsPurchased ?? 0}
              </Text>
              <Text bold fontSize={16} color='green'>
                Checked: {NFT.totalNFTsChecked ?? 0}
              </Text>
            </HStack>

            {user?.role === UserRole.Collector && (
              <Button
                variant='outlined'
                disabled={isLoading}
                onPress={() => buyNFT(NFT.id)}
              >
                Buy NFT
              </Button>
            )}

            <Text fontSize={13} color='gray'>
              {new Date(NFT.date).toLocaleString()}
            </Text>
          </VStack>
        )}
      />
    </VStack>
  );
}

const HeaderRight = () => {
  return (
    <TabBarIcon
      size={32}
      name="add-circle-outline"
      onPress={() => router.push('/(nft-creation)/new')}
    />
  );
};
