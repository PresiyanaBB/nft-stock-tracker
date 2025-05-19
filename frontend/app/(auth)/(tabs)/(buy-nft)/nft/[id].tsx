import { Button } from '@/components/Button';
import DateTimePicker from '@/components/DateTimePicker';
import { Input } from '@/components/Input';
import { Text } from '@/components/Text';
import { VStack } from '@/components/VStack';
import { TabBarIcon } from '@/components/navigation/TabBarIcon';
import { NFTService } from '@/services/NFTs';
import { NFT } from '@/types/NFT';
import { useFocusEffect } from '@react-navigation/native';
import { useLocalSearchParams, useNavigation, router } from 'expo-router';
import { useCallback, useEffect, useState } from 'react';
import { Alert } from 'react-native';
import { validate as validateUUID } from 'uuid';

export default function NFTDetailsScreen() {
  const navigation = useNavigation();
  const { id } = useLocalSearchParams();

  const [isSubmitting, setIsSubmitting] = useState(false);
  const [NFTData, setNFTData] = useState<NFT | null>(null);

  function updateField(field: keyof NFT, value: string | Date) {
    setNFTData((prev) => ({
      ...prev!,
      [field]: value,
    }));
  }

  const onDelete = useCallback(async () => {
    if (!NFTData) return;
    if (!validateUUID(id as string)) {
      Alert.alert("Error", "Invalid UUID");
      return;
    }

    try {
      Alert.alert("Delete NFT", "Are you sure you want to delete this NFT?", [
        { text: "Cancel" },
        {
          text: "Delete",
          onPress: async () => {
            await NFTService.deleteNFT(id as string);
            router.back();
          }
        },
      ]);
    } catch (error) {
      Alert.alert("Error", "Failed to delete NFT");
    }
  }, [NFTData, id]);

  async function onSubmitChanges() {
    if (!NFTData) return;
    if (!validateUUID(id as string)) {
      Alert.alert("Error", "Invalid UUID");
      return;
    }

    try {
      setIsSubmitting(true);
      await NFTService.updateNFT(id as string, {
        name: NFTData.name,
        token_uri: NFTData.token_uri,
        creator: NFTData.creator,
        price: NFTData.price,
        date: NFTData.date,
      });
      router.back();
    } catch (error) {
      Alert.alert("Error", "Failed to update NFT");
    } finally {
      setIsSubmitting(false);
    }
  }

  const fetchNFT = async () => {
    if (!validateUUID(id as string)) {
      Alert.alert("Error", "Invalid UUID");
      router.back();
      return;
    }

    try {
      const response = await NFTService.getNFT(id as string);
      setNFTData(response.data);
    } catch (error) {
      Alert.alert("Error", "Failed to fetch NFT");
      router.back();
    }
  };

  useFocusEffect(
    useCallback(() => {
      fetchNFT();
    }, [id])
  );

  useEffect(() => {
    navigation.setOptions({
      headerTitle: NFTData?.name || "NFT Details",
      headerRight: () => <HeaderRight onDelete={onDelete} />,
    });
  }, [navigation, NFTData, onDelete]);

  return (
    <VStack m={20} flex={1} gap={30}>
      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">Name</Text>
        <Input
          value={NFTData?.name || ""}
          onChangeText={(value) => updateField("name", value)}
          placeholder="Name"
          placeholderTextColor="darkgray"
          h={48}
          p={14}
        />
      </VStack>

      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">Token URI</Text>
        <Input
          value={NFTData?.token_uri || ""}
          onChangeText={(value) => updateField("token_uri", value)}
          placeholder="Token URI"
          placeholderTextColor="darkgray"
          h={48}
          p={14}
        />
      </VStack>

      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">Creator</Text>
        <Input
          value={NFTData?.creator || ""}
          onChangeText={(value) => updateField("creator", value)}
          placeholder="Creator"
          placeholderTextColor="darkgray"
          h={48}
          p={14}
        />
      </VStack>

      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">Price</Text>
        <Input
          value={NFTData?.price?.toString() || ""}
          onChangeText={(value) => updateField("price", value)}
          placeholder="Price"
          placeholderTextColor="darkgray"
          h={48}
          p={14}
          keyboardType="decimal-pad"
        />
      </VStack>

      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">Date</Text>
        <DateTimePicker
          onChange={(date) => updateField('date', date || new Date())}
          currentDate={NFTData?.date ? new Date(NFTData.date) : new Date()}
        />
      </VStack>

      <Button
        mt={"auto"}
        isLoading={isSubmitting}
        disabled={isSubmitting}
        onPress={onSubmitChanges}
      >
        Save Changes
      </Button>
    </VStack>
  );
}

const HeaderRight = ({ onDelete }: { onDelete: VoidFunction }) => {
  return (
    <TabBarIcon size={30} name="trash" onPress={onDelete} />
  );
};
