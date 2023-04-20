export interface WebsocketMessage<T> {
	event: string
	payload: T
}

export interface LabState {
	dbValid: null
	taskTitle: string
	taskDescription: string
	taskTags: TaskTag[]
	query: string
	queryPassed: boolean
	queryError: null
	result: Result
}

export interface Result {
	expectedHeader: string[]
	expectedRows: Array<string[]>
	actualHeader: string[]
	actualRows: Array<string[]>
}

export interface TaskTag {
	key: string
	value: string
}
