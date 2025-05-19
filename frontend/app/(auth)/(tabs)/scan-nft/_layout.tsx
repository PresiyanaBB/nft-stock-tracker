import { Button } from '@/components/Button';
import { Text } from '@/components/Text';
import { VStack } from '@/components/VStack';
import { userNFTService } from '@/services/userNFTs';
import { BarcodeScanningResult, CameraView, useCameraPermissions } from 'expo-camera';
import { useState } from 'react';
import { ActivityIndicator, Alert, Vibration } from 'react-native';
import { validate as validateUUID } from 'uuid';

export default function ScanUserNFTScreen() {
  const [permission, requestPermission] = useCameraPermissions();
  const [scanningEnabled, setScanningEnabled] = useState(true);
  const [isLoading, setIsLoading] = useState(false);

  if (!permission) {
    return (
      <VStack flex={1} justifyContent="center" alignItems="center">
        <ActivityIndicator size="large" />
      </VStack>
    );
  }

  if (!permission.granted) {
    return (
      <VStack gap={20} flex={1} justifyContent="center" alignItems="center">
        <Text>Camera access is required to scan the NFT.</Text>
        <Button onPress={requestPermission}>Allow Camera Access</Button>
      </VStack>
    );
  }

  async function onBarcodeScanned({ data }: BarcodeScanningResult) {
    if (!scanningEnabled) return;

    try {
      Vibration.vibrate();
      setScanningEnabled(false);
      setIsLoading(true);
      console.log("Scanned data:", data);
      // Parsing the scanned data
      const [userNFT, owner] = data.split(",");
      const UserNFTId = userNFT.split(":")[1].trim();
      const OwnerId = owner.split(":")[1].trim();

      // Validating the NFT
      await userNFTService.validateUserNFT(UserNFTId, OwnerId);
      const nftId = (await userNFTService.getAll()).data.find(userNFT => userNFT.id === UserNFTId)?.nft_id;
      Alert.alert('Success', "NFT validated successfully.", [
        { text: 'Ok', onPress: () => setScanningEnabled(true) },
      ]);

      const deleted = (await userNFTService.getAll()).data.filter(userNFT => {
        return userNFT.id !== UserNFTId && userNFT.nft_id === nftId;

      });

      // Delete all other user NFTs so that they can't be scanned again
      deleted.forEach(async (userNFT) => {
        await userNFTService.deleteUserNFT(userNFT.id);
      });
    } catch (error: any) {
      Alert.alert('Error', error.message || "Failed to validate NFT. Please try again.");
      setScanningEnabled(true);
    }
  }

  return (
    <VStack flex={1}>
      <CameraView
        style={{ flex: 1 }}
        facing="back"
        onBarcodeScanned={scanningEnabled ? onBarcodeScanned : undefined}
        barcodeScannerSettings={{
          barcodeTypes: ["qr"],
        }}
      />

      {isLoading && (
        <VStack
          position="absolute"
          top={0}
          bottom={0}
          left={0}
          right={0}
          justifyContent="center"
          alignItems="center"
          backgroundColor="rgba(0,0,0,0.5)"
        >
          <ActivityIndicator size="large" color="white" />
          <Text mt={10} color="white">Validating NFT...</Text>
        </VStack>
      )}
    </VStack>
  );
}
