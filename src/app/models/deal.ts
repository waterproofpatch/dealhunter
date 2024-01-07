import { Location } from "./location";

export interface Deal {
	Location: Location
	RetailPrice: number
	ActualPrice: number
	StoreName: string
	ItemName: string

}