import { Box, Typography } from '@mui/material'
import React from 'react'
import EnrollmentItem from './EnrollmentItem'

interface EnrollmentDrawerProps {
	enrollSubmit: () => void
}

const EnrollmentDrawer: React.FC<EnrollmentDrawerProps> = ({
	enrollSubmit,
}) => {
	return (
		<Box
			sx={{
				width: '35vw',
				backgroundColor: '#FAFAFA',
				height: '100%',
				py: 8,
				px: 4,
			}}
		>
			<Typography fontSize={24} px={2} mb={5}>
				Available Labs
			</Typography>
			<EnrollmentItem enrollSubmit={enrollSubmit} />
		</Box>
	)
}

export default EnrollmentDrawer
