import { Button } from '@/components/Button';
import { Input } from '@/components/Input';
import { Text } from '@/components/Text';
import { VStack } from '@/components/VStack';
import { NFTService } from '@/services/NFTs';
import { useNavigation, router } from 'expo-router';
import React, { useEffect, useState } from 'react';
import { Alert, Image, TouchableOpacity } from 'react-native';
import * as ImagePicker from 'expo-image-picker';

export default function NewNFT() {
  const navigation = useNavigation();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [tokenUri, setTokenUri] = useState('');
  const [name, setName] = useState('');
  const [creator, setCreator] = useState('');
  const [price, setPrice] = useState('');
  const [image, setImage] = useState<string | null>(null);
  const [imageBase64, setImageBase64] = useState<string | null>(null);


  async function onSubmit() {
    if (!name || !creator || !price || !imageBase64) {
      Alert.alert("Error", "Please fill out all fields and select an image.");
      return;
    }

    try {
      setIsSubmitting(true);
      await NFTService.createNFT(tokenUri, name, creator, price, imageBase64);
      Alert.alert("Success", "NFT created successfully");
      router.back();
    } catch (error) {
      console.error(error);
      Alert.alert("Error", "Failed to create NFT");
    } finally {
      setIsSubmitting(false);
    }
  }

  async function pickImage() {
    const { status } = await ImagePicker.requestMediaLibraryPermissionsAsync();
    if (status !== 'granted') {
      Alert.alert("Permission Required", "We need your permission to access the image library.");
      return;
    }

    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.Images,
      base64: true,
      quality: 1,
    });

    if (!result.canceled) {
      setImage(result.assets[0].uri);
      setImageBase64(result.assets[0].base64 ?? null);
    }
  }

  useEffect(() => {
    navigation.setOptions({ headerTitle: "New NFT" });
  }, []);

  return (
    <VStack m={20} flex={1} gap={20}>
      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">TokenURI</Text>
        <Input
          value={tokenUri}
          onChangeText={setTokenUri}
          placeholder="TokenURI"
          placeholderTextColor="darkgray"
          h={48}
          p={14}
        />
      </VStack>

      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">Name</Text>
        <Input
          value={name}
          onChangeText={setName}
          placeholder="Name"
          placeholderTextColor="darkgray"
          h={48}
          p={14}
        />
      </VStack>

      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">Creator</Text>
        <Input
          value={creator}
          onChangeText={setCreator}
          placeholder="Creator"
          placeholderTextColor="darkgray"
          h={48}
          p={14}
        />
      </VStack>

      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">Price</Text>
        <Input
          value={price}
          onChangeText={setPrice}
          placeholder="Price (e.g., 10.00)"
          placeholderTextColor="darkgray"
          h={48}
          p={14}
          keyboardType="decimal-pad"
        />
      </VStack>

      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">Image</Text>
        <TouchableOpacity onPress={pickImage} style={{ alignItems: 'center' }}>
          {image ? (
            <Image source={{ uri: image }} style={{ width: 150, height: 150, borderRadius: 10 }} />
          ) : (
            <Text style={{ color: 'dodgerblue', fontSize: 16 }}>Select Image</Text>
          )}
        </TouchableOpacity>
      </VStack>

      <Button
        mt="auto"
        isLoading={isSubmitting}
        disabled={isSubmitting}
        onPress={onSubmit}
      >
        Save
      </Button>
    </VStack>
  );
}
