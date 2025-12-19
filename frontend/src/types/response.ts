export class APIResponse<T> {
	success: boolean;
	data: T | string;

	constructor(success: boolean, data: T | string) {
		this.success = success;
		this.data = data;
	}

	payload(): T {
		if (!this.success) {
			throw new TypeError('cannot get payload of failed response');
		}
		return this.data as T;
	}

	error(): string {
		if (this.success) {
			throw new TypeError('cannot get error of successful response');
		}
		return this.data as string;
	}
}
