import { Location } from "./location";

export interface Deal {
	location: Location
	retailPrice: number
	actualPrice: number
	storeName: string
	itemName: string

}