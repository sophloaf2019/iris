import { inject } from '@angular/core';
import { HTTPWrapper } from '../../services/http';

export abstract class CrudService<T extends Object> {
	protected http: HTTPWrapper = inject(HTTPWrapper);
	protected constructor(protected base: string) {}

	async list(): Promise<T[]> {
		return this.http.get<T[]>(this.base);
	}

	async getByID(id: number): Promise<T> {
		return this.http.get<T>(`${this.base}?id=${id}`);
	}

	async getBySlug(slug: string): Promise<T> {
		return this.http.get<T>(`${this.base}?slug=${slug}`);
	}

	async create(item: T): Promise<T> {
		return this.http.post<T>(this.base, item);
	}

	async update(id: number, item: T): Promise<T> {
		return this.http.put<T>(`${this.base}/${id}`, item);
	}

	async delete(id: number) {
		await this.http.delete<void>(`${this.base}/${id}`);
	}
}
