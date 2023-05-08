import { Box, Stack, Typography } from '@mui/material'
import React, { useEffect } from 'react'
import EnrollmentItem from './EnrollmentItem'
import axios from 'axios'
import { BasedResponse } from '../types/APIs/basedResponse'
import { LabAPI, LabDetailAPI } from '../types/APIs/Profile/lab'

interface EnrollmentDrawerProps {
	enrollSubmit: (a0: number) => void
}

const EnrollmentDrawer: React.FC<EnrollmentDrawerProps> = ({
	enrollSubmit,
}) => {
	const [labs, setLabs] = React.useState<LabDetailAPI[]>([])

	useEffect(() => {
		axios
			.get<BasedResponse<LabAPI>>('/api/profile/labs')
			.then((response) => {
				setLabs(response.data.data.labs)
			})
			.catch((error) => {
				alert(error.response.data.message)
			})
	}, [])
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
			<Stack gap={2}>
				{labs.map((lab) => (
					<EnrollmentItem
						enrollSubmit={enrollSubmit}
						lab={lab}
						key={lab.id}
					/>
				))}
			</Stack>
		</Box>
	)
}

export default EnrollmentDrawer
