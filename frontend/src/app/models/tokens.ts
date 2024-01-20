import { jwtDecode } from 'jwt-decode';

export class JwtAccessToken {
	accessToken: string;
	id: number = 0;
	email: string = "";
	exp: any;

	constructor(accessToken: string) {
		this.accessToken = accessToken;
		const decodedToken: any = jwtDecode(this.accessToken);
		this.id = decodedToken['id'];
		this.email = decodedToken['email'];
		this.exp = decodedToken['exp'];
	}

	public isExpired(): boolean {
		const currentTime = Math.floor(Date.now() / 1000); // Get the current time in seconds
		return currentTime > this.exp;
	}

	public secondsUntilExpiration(): number {
		const currentTime = Math.floor(Date.now() / 1000); // Get the current time in seconds
		return this.exp - currentTime;
	}
}
