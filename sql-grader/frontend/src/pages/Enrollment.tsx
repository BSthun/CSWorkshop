import { Add } from '@mui/icons-material'
import { Box, Button, SwipeableDrawer, Typography } from '@mui/material'
import React, { useState } from 'react'
import EnrollmentDrawer from '../components/EnrollmentDrawer'
import EnrollmentBox from '../components/EnrollmentBox'
import { EnrollStateAPI } from '../types/APIs/Profile/enroll_state'
import axios from 'axios'

const Enrollment = () => {
	const [_, setEnrollState] = useState<EnrollStateAPI>()

	const enrollSubmit = async () => {
		const response = await axios.post<EnrollStateAPI>(
			'/api/profile/enroll',
			{
				lab_id: 1,
				source: 'blank',
			}
		)
		setEnrollState(response.data)
	}

	const [openMenu, setOpenMenu] = React.useState(false)
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
				<EnrollmentBox
					dbName="lab_spotify01_11cnel"
					dbValid={true}
					labName="Spotiy Test"
					enrollAt={new Date()}
				/>
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
