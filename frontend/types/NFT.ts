import { ApiResponse } from "./api";

export type NFTResponse = ApiResponse<NFT>;
export type NFTListResponse = ApiResponse<NFT[]>;

export type NFT = {
  id: string; // UUID as string
  tokenUri: string;
  name: string;
  creator: string;
  price: number; // Decimal as number
  image: string; // Base64 image as string
  collected: boolean;
  createdAt: string;
  updatedAt: string;
};
