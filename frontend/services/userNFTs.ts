import { UserNFT, UserNFTListResponse, UserNFTResponse } from "@/types/userNFTs";
import { Api } from "./api";
import { ApiResponse } from "@/types/api";

async function createUserNFT(nftId: string): Promise<UserNFTResponse> {
  return Api.post("/userNFT", { nftId });
}

async function getUserNFT(id: string): Promise<ApiResponse<{ userNFT: UserNFT, qrcode: string }>> {
  return Api.get(`/userNFT/${id}`);
}

async function getAll(): Promise<UserNFTListResponse> {
  return Api.get("/userNFT");
}

async function validateUserNFT(userNFTId: string, ownerId: string): Promise<UserNFTResponse> {
  return Api.post("/userNFT/validate", { userNFTId, ownerId });
}

const userNFTService = {
  createUserNFT,
  getUserNFT,
  getAll,
  validateUserNFT,
}

export { userNFTService }
