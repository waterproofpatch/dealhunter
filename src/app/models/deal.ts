import { Location } from "./location";

export interface Deal {
	ID: number
	CreatedAt: string
	Location: Location
	RetailPrice: number
	ActualPrice: number
	StoreName: string
	ItemName: string
	Upvotes: number
	LastUpvoteTime: string

}