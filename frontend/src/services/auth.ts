import { inject, Injectable } from '@angular/core';
import { Action, Content, User } from '../types/auth/user';
import { HTTPWrapper } from './http';
import { Store } from './store';

// The AuthService is responsible for interacting with the
// authentication and authorization systems of the backend.
@Injectable({ providedIn: 'root' })
export class AuthService {
	// Member http represents the HTTPWrapper which
	// the AuthService will use.
	private http: HTTPWrapper = inject(HTTPWrapper);

	// Member auth_base represents the base URL for auth-specific requests.
	private authBase: string = 'auth';
	// Member user_base represents the base URL for user-specific requests.
	private userBase: string = 'user';

	// Member currentUser represents who is currently logged in.
	private userKey: string = 'iris_user';

	private storage: Store = inject(Store);

	// Getter user() returns the current user.
	get user(): User | null {
		return this.storage.get<User>(this.userKey);
	}

	// Setter user() sets the current user via local storage.
	// If null is passed, the user entry is deleted.
	set user(user: User | null) {
		this.storage.store<User>(this.userKey, user);
	}

	// Function clear_user() calls the setter for user with null.
	clearUser() {
		this.user = null;
	}

	// Member tokenKey represents the local storage key for use
	// with authorizing on the backend. Local storage will store
	// the session token under this key.
	private tokenKey = 'iris_token';

	// Getter token() returns the current token from local storage.
	get token(): string | null {
		return this.storage.get<string>(this.tokenKey);
	}

	// Setter token() sets the current token. If an empty string is passed,
	// the token is removed from local storage.
	set token(token: string) {
		this.storage.store<string>(this.tokenKey, token);
	}

	// Function clear_token() calls the setter for token with an empty string.
	clearToken() {
		this.token = '';
	}

	// Function login() takes in a username and password. It attempts
	// to log the user in, and -- if successful -- stores the token
	// and user.
	async login(username: string, password: string): Promise<string> {
		this.token = await this.http.post<string>(`${this.authBase}/login`, {
			username,
			password,
		});
		this.user = await this.getUserByUsername(username);
		return this.token;
	}

	// Function logout() calls the logout route from the backend and clears the token,
	// regardless of if the logout was successful or not (but why would it ever fail?).
	async logout() {
		await this.http.post<null>(`${this.authBase}/logout`, {});
		this.clearToken();
		this.clearUser();
	}

	// Function issueUser() takes a username, password, and clearance level, and returns the API response.
	async issueUser(username: string, password: string, clearance: number): Promise<User> {
		return await this.http.post<User>(`${this.authBase}/issue_user`, {
			username,
			password,
			clearance,
		});
	}

	// Function resetPassword() takes a userID and then their old and new password.
	// Note that old password may be empty if you're an administrator.
	async resetPassword(user_id: number, old_password: string, new_password: string) {
		await this.http.post<null>(`${this.authBase}/reset_password`, {
			user_id,
			old_password,
			new_password,
		});
	}

	// Function get_user_by_id() takes an ID number and returns the results of the query.
	async getUserByID(user_id: number): Promise<User> {
		return await this.http.get<User>(`${this.userBase}?id=${user_id}`);
	}

	// Function getUserByUsername() takes a username and returns the results of the query.
	async getUserByUsername(username: string): Promise<User> {
		return await this.http.get<User>(`${this.userBase}?username=${username}`);
	}

	// Function hi() attempts to say hi. It 401s if the user is not authenticated,
	// and 200s if it does.
	async hi(): Promise<string> {
		return await this.http.get<string>(`${this.authBase}/hi`);
	}

	async userCan(
		userID: number,
		action: Action,
		content: Content,
		isOwner: boolean,
	): Promise<boolean> {
		return await this.http.post<boolean>(`${this.userBase}/can`, {
			userID,
			action,
			content,
			isOwner,
		});
	}
}
