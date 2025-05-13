import { ApiResponse } from "./api";
import { validate as validateUUID, v4 as uuidv4 } from 'uuid';

export type NFTResponse = ApiResponse<NFT>;
export type NFTListResponse = ApiResponse<NFT[]>;

export type NFT = {
  id: string; // UUID as string
  tokenUri: string;
  name: string;
  creator: string;
  price: number; // Decimal as number
  image: string; // Base64 image as string
  totalNFTsPurchased: number;
  totalNFTsChecked: number;
  createdAt: string;
  updatedAt: string;
};
