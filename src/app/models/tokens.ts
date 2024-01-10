
import { jwtDecode } from 'jwt-decode';


export class JwtAccessToken {
	accessToken: string;
	id: number = 0;
	email: string = ""

	constructor(accessToken: string) {
		this.accessToken = accessToken;
		const decodedToken: any = jwtDecode(this.accessToken);
		this.id = decodedToken['id']
		this.email = decodedToken['email']
	}
}
