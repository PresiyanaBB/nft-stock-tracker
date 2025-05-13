import { NFTResponse, NFTListResponse } from "@/types/NFT";
import { Api } from "./api";

async function createNFT(name: string, location: string, date: string): Promise<NFTResponse> {
  return Api.post("/nft", { name, location, date });
}

async function getNFT(id: string): Promise<NFTResponse> {
  return Api.get(`/nft/${id}`);
}

async function getAll(): Promise<NFTListResponse> {
  return Api.get("/nft");
}

async function updateNFT(id: string, tokenUri:string, name: string, creator: string, price: number, image: string): Promise<NFTResponse> {
  return Api.put(`/nft/${id}`, { name, location, date });
} 

async function deleteNFT(id: string): Promise<NFTResponse> {
  return Api.delete(`/nft/${id}`);
}

const NFTService = {
  createNFT,
  getNFT,
  getAll,
  updateNFT,
  deleteNFT,
};

export { NFTService };
