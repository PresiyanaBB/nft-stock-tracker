import { UserNFT, UserNFTListResponse, UserNFTResponse } from "@/types/userNFT";
import { NFTService } from "@/services/NFTs";
import { Api } from "./api";
import { ApiResponse } from "@/types/api";

async function createUserNFT(nftId: string): Promise<UserNFTResponse> {
  return Api.post("/userNFT", { nft_id: nftId });
}

async function getUserNFT(userNFT_id: string): Promise<ApiResponse<{ userNFT: UserNFT, qrcode: string }>> {
  const response = await Api.get(`/userNFT/${userNFT_id}`);
  return { data: response.data, message: "ok", status: "success" };
}

async function getAll(): Promise<UserNFTListResponse> {
  const response = await Api.get("/userNFT");

  response.data.forEach(async (userNFT: UserNFT) => {
    try {
      const nftResponse = await NFTService.getNFT(userNFT.nft_id);
      userNFT.nft = nftResponse.data;
    } catch (error) {
      console.error(`Failed to fetch NFT for UserNFT ${userNFT.id}`, error);
    }
  });

  return { data: response.data, message: "ok", status: "success" };
}

async function validateUserNFT(UserNFTId: string, OwnerId: string): Promise<UserNFTResponse> {
  return Api.post("/userNFT/validate", { UserNFTId, OwnerId });
}

const userNFTService = {
  createUserNFT,
  getUserNFT,
  getAll,
  validateUserNFT,
}

export { userNFTService }
