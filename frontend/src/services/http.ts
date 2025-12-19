import { inject, Injectable } from '@angular/core';
import { APIResponse } from '../types/response';
import { AuthService } from './auth';
import { Store } from './store';
import { ActivatedRoute, Route, Router } from '@angular/router';

@Injectable({ providedIn: 'root' })
export class HTTPWrapper {
	private root: string = 'http://localhost:8080/api/';
	private storage: Store = inject(Store);
	private router: Router = inject(Router);

	// Internal wrapper for fetch requests
	private async request<T>(
		method: 'GET' | 'POST' | 'PUT' | 'DELETE',
		url: string,
		body?: Object,
	): Promise<T> {
		console.log(this.root + url);
		let token = this.storage.get<string>('iris_token');
		const response = await fetch(this.root + url, {
			method,
			headers: {
				'Content-Type': 'application/json',
				Authorization: token ? `Bearer ${token}` : '',
			},
			body: body ? JSON.stringify(body) : undefined,
		});
		if (!response.ok) {
			if (
				response.status === 401 &&
				this.router.url !== '' &&
				this.router.url !== '/' &&
				!this.router.url.startsWith('/?next')
			) {
				await this.router.navigate(['/'], {
					queryParams: { next: this.router.url },
				});
			}
			throw new Error(`HTTP error! status: ${response.status}`);
		}
		const data: any = await response.json();
		return data['data'] as T;
	}

	// Public methods
	async get<T>(url: string): Promise<T> {
		return this.request<T>('GET', url);
	}

	post<T>(url: string, body: Object): Promise<T> {
		return this.request<T>('POST', url, body);
	}

	put<T>(url: string, body: Object): Promise<T> {
		return this.request<T>('PUT', url, body);
	}

	delete<T>(url: string): Promise<T> {
		return this.request<T>('DELETE', url);
	}
}
