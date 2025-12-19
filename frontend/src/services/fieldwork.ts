import { Injectable } from '@angular/core';
import { GOI } from '../types/fieldwork/goi';
import { CrudService } from '../types/fieldwork/crud';
import { Mission } from '../types/fieldwork/mission';
import { Debrief } from '../types/fieldwork/debrief';

@Injectable({ providedIn: 'root' })
export class GOIService extends CrudService<GOI> {
	constructor() {
		super('http://localhost:8080/api/goi');
	}
}

@Injectable({ providedIn: 'root' })
export class MissionService extends CrudService<Mission> {
	constructor() {
		super('http://localhost:8080/api/mission');
	}
}

@Injectable({ providedIn: 'root' })
export class DebriefService extends CrudService<Debrief> {
	constructor() {
		super('http://localhost:8080/api/debrief');
	}
}
