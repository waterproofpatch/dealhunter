
import * as jwt_decode from 'jwt-decode';

export class JwtAccessToken {
	accessToken: string;

	constructor(accessToken: string) {
		this.accessToken = accessToken;
	}

	decodeToken() {
		try {
			const decodedToken = jwt_decode(this.accessToken);
			console.log(decodedToken);
			return decodedToken;
		} catch (error) {
			console.error("Error decoding token", error);
			return null;
		}
	}
}
