export interface LabAPI {
	labs: LabDetailAPI[]
}

export interface LabDetailAPI {
	id: number
	code: string
	name: string
	description: string
}
