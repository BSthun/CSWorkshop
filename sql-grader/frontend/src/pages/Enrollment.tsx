import { Add } from '@mui/icons-material'
import { Box, Button, SwipeableDrawer, Typography } from '@mui/material'
import React, { useEffect, useState } from 'react'
import EnrollmentDrawer from '../components/EnrollmentDrawer'
import EnrollmentBox from '../components/EnrollmentBox'
import { EnrollStateAPI } from '../types/APIs/Profile/enroll_state'
import axios from 'axios'
import { EnrollmentsAPI } from '../types/APIs/Profile/enrollment'
import { BasedResponse } from '../types/APIs/basedResponse'

const Enrollment = () => {
	const [_, setEnrollState] = useState<EnrollStateAPI>()
	const [enrollments, setEnrollments] = useState<EnrollmentsAPI[]>()
	const [openMenu, setOpenMenu] = React.useState(false)

	const enrollSubmit = async () => {
		try {
			const response = await axios.post<EnrollStateAPI>(
				'/api/profile/enroll',
				{
					lab_id: 1,
					source: 'blank',
				}
			)
			setEnrollState(response.data)

			if (response.data.success) {
				await getEnrollments()
			} else {
				alert(response.data.message)
			}
			setOpenMenu(false)
		} catch {
			alert('An error occured')
			setOpenMenu(false)
		}
	}

	const getEnrollments = async () => {
		try {
			const response = await axios.get<BasedResponse<EnrollmentsAPI[]>>(
				'/api/profile/enrollments'
			)
			setEnrollments(response.data.data)
		} catch {
			alert('An error occured')
		}
	}

	useEffect(() => {
		getEnrollments()
	}, [])

	return (
		<Box p="80px">
			<Box
				sx={{
					display: 'flex',
					justifyContent: 'space-between',
					mb: 3,
				}}
			>
				<Typography fontSize={36}>My Enrollments</Typography>
				<Button
					startIcon={<Add />}
					variant="contained"
					sx={{
						backgroundColor: '#FCAC4E',
						textTransform: 'none',
						width: '150px',
						height: '50px',
						':hover': {
							bgcolor: '#FCAC4E',
							color: 'white',
						},
						fontSize: 18,
						borderRadius: 2,
					}}
					onClick={() => setOpenMenu(true)}
				>
					Enroll
				</Button>
			</Box>
			<Box>
				{enrollments?.map((item, index) => (
					<EnrollmentBox
						dbName={item.dbName}
						dbValid={item.dbValid}
						labName={item.labName}
						enrollAt={item.enrolledAt}
						key={index}
					/>
				))}
			</Box>
			<SwipeableDrawer
				anchor="right"
				open={openMenu}
				onClose={() => setOpenMenu(false)}
				onOpen={() => setOpenMenu(true)}
			>
				<EnrollmentDrawer enrollSubmit={enrollSubmit} />
			</SwipeableDrawer>
		</Box>
	)
}

export default Enrollment
