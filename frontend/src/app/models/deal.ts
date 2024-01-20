import { Location } from "./location";
import { User } from "./user";

export interface Deal {
	ID: number
	CreatedAt: string
	Location: Location
	User: User,
	RetailPrice: number
	ActualPrice: number
	StoreName: string
	ItemName: string
	Upvotes: number
	LastUpvoteTime: string

}