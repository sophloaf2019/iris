import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Post } from '../types/social/post';
import { APIResponse } from '../types/response';
import { Observable } from 'rxjs';
import { HTTPWrapper } from './http';
import { Comment } from '../types/social/comment';

@Injectable({ providedIn: 'root' })
export abstract class PostService {
	protected http: HTTPWrapper = inject(HTTPWrapper);
	private base: string = 'posts';

	list(): Promise<Post[]> {
		return this.http.get<Post[]>(this.base);
	}

	get(id: number): Promise<Post> {
		return this.http.get<Post>(`${this.base}/${id}`);
	}

	create(item: Post): Promise<Post> {
		return this.http.post<Post>(this.base, item);
	}

	update(id: number, item: Post): Promise<Post> {
		return this.http.put<Post>(`${this.base}/${id}`, item);
	}

	delete(id: number): Promise<void> {
		return this.http.delete<void>(`${this.base}/${id}`);
	}
}

@Injectable({ providedIn: 'root' })
export abstract class CommentService {
	protected http: HTTPWrapper = inject(HTTPWrapper);
	private base: string = 'comments';

	list(): Promise<Comment[]> {
		return this.http.get<Comment[]>(this.base);
	}

	getByID(id: number): Promise<Comment> {
		return this.http.get<Comment>(`${this.base}?id=${id}`);
	}

	getForParent(contentType: string, parentID: number): Promise<Comment[]> {
		return this.http.get<Comment[]>(
			`${this.base}?content_type=${contentType}&parent_id=${parentID}`,
		);
	}

	create(item: Comment): Promise<Comment> {
		return this.http.post<Comment>(this.base, item);
	}

	update(id: number, item: Comment): Promise<Comment> {
		return this.http.put<Comment>(`${this.base}/${id}`, item);
	}

	delete(id: number): Promise<void> {
		return this.http.delete<void>(`${this.base}/${id}`);
	}
}

