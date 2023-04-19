export interface LabAPI {
	labs: LabDetailAPI[]
}

interface LabDetailAPI {
	id: number
	code: string
	name: string
	description: string
}
