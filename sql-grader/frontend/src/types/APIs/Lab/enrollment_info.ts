export interface EnrollmentInfoAPI {
	enrollmentId: number
	enrolledAt: Date
	labName: string
	dbName: string
	dbValid: false
	dbHost: string
	dbPort: string
	dbUsername: string
	dbPassword: string
	tasks: EnrollmentInfoTask[]
	token: string
}

export interface EnrollmentInfoTask {
	id: number
	title: string
	passed: boolean
}
