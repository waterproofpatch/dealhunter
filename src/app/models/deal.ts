import { Location } from "./location";

export interface Deal {
	CreatedAt: string
	Location: Location
	RetailPrice: number
	ActualPrice: number
	StoreName: string
	ItemName: string

}