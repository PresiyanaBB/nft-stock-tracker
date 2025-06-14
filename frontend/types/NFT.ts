import { ApiResponse } from "./api";

export type NFTResponse = ApiResponse<NFT>;
export type NFTListResponse = ApiResponse<NFT[]>;

export type NFT = {
  id: string;
  tokenUri: string;
  name: string;
  creator: string;
  price: number;
  image: string;
  collected: boolean;
  createdAt: string;
  updatedAt: string;
};
