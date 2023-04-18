export interface EnrollmentInfoAPI {
    enrollmentId: number
    enrolledAt: Date
    labName: string
    dbName: number
    dbValid: false
    dbHost: string
    dbPort: string
    dbUsername: string
    dbPassword: string
    tasks: EnrollmentInfoTask[]
}

interface EnrollmentInfoTask {
    id: number
    title: string
}