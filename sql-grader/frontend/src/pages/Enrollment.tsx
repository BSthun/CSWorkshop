import { Add } from '@mui/icons-material'
import { Box, Button, Typography } from '@mui/material'
import React from 'react'

const Enrollment = () => {
	return (
		<Box p="80px">
			<Box
				sx={{
					display: 'flex',
					justifyContent: 'space-between',
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
						height: '53px',
					}}
				>
					Enroll
				</Button>
			</Box>
		</Box>
	)
}

export default Enrollment
