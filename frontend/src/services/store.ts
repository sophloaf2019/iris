import { Injectable } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class Store {
	store<T>(key: string, object: T | null) {
		if (object != null) {
			localStorage.setItem(key, JSON.stringify(object));
		} else {
			localStorage.removeItem(key);
		}
	}

	get<T>(key: string): T | null {
		let obj = localStorage.getItem(key);
		if (obj == null) {
			return null;
		}
		try {
			return JSON.parse(obj) as T;
		} catch {
			return obj as T;
		}
	}
}
