import { Box, Button, Typography } from '@mui/material'
import React from 'react'
import { LabDetailAPI } from '../types/APIs/Profile/lab'

interface EnrollmentItemProps {
	lab: LabDetailAPI
	enrollSubmit: (a0: number) => void
}

const EnrollmentItem: React.FC<EnrollmentItemProps> = ({
	lab,
	enrollSubmit,
}) => {
	return (
		<Box
			sx={{
				display: 'flex',
				flexDirection: 'column',
				px: 4,
				py: 2.5,
				boxShadow: '0px 2px 12px rgba(0, 0, 0, 0.25)',
				borderRadius: 2.5,
			}}
		>
			<Typography fontSize={18}>{lab.code}</Typography>
			<Typography fontSize={20} fontWeight={700}>
				{lab.name}
			</Typography>
			<Typography>{lab.description}</Typography>
			<Button
				variant="outlined"
				sx={{
					mt: 3,
					width: '144px',
					height: '44px',
					backgroundColor: '#E1F6FB',
					border: 'none',
					textTransform: 'none',
					fontSize: 16,
					':hover': {
						border: 'none',
					},
					alignSelf: 'end',
					borderRadius: 2,
				}}
				onClick={() => enrollSubmit(lab.id)}
			>
				Enroll
			</Button>
		</Box>
	)
}

export default EnrollmentItem
