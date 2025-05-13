import { ApiResponse } from "./api";
import { NFT } from "./NFT";

export type UserNFTResponse = ApiResponse<UserNFT>;
export type UserNFTListResponse = ApiResponse<UserNFT[]>;

export type UserNFT = {
  id: string;
  nftId: string;
  userId: string;
  NFT: NFT;
  collected: boolean;
  createdAt: string;
  updatedAt: string;
};
