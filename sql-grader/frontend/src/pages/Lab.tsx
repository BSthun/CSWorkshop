import {
	Autorenew,
	BatchPrediction,
	FormatListBulleted,
} from '@mui/icons-material'
import {
	Box,
	Button,
	CircularProgress,
	Dialog,
	Divider,
	IconButton,
	Table,
	TableBody,
	TableCell,
	TableContainer,
	TableHead,
	TableRow,
	TextField,
	Typography,
} from '@mui/material'
import React from 'react'
import TaskItemList from '../components/TaskItemList'
import axios from 'axios'
import { BasedResponse } from '../types/APIs/basedResponse'
import { EnrollmentInfoAPI } from '../types/APIs/Lab/enrollment_info'
import { useParams } from 'react-router-dom'

interface DbInfo {
	label: string
	value: string
}

const taskTags = [
	{
		name: 'Level',
		value: 'Lv1',
	},
	{
		name: 'Difficulty',
		value: 'Easy',
	},
	{
		name: 'Tag 3',
		value: 'Hello',
	},
]

const Lab = () => {
	const [isTask, setIsTask] = React.useState(false)
	const [selectedTask, setSelectedTask] = React.useState(0)
	const [isLoading, setIsLoading] = React.useState(false)
	const [isDialogOpen, setIsDialogOpen] = React.useState(false)
	const [log, setLog] = React.useState(
		"Lorem Ipsum is simply dummy text of the printing and typesetting industry. <br /> Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum. Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum"
	)
	const [enrollmentInfo, setEnrollmentInfo] =
		React.useState<EnrollmentInfoAPI>()
	const [dbInfo, setDbInfo] = React.useState<DbInfo[]>([])
	const params = useParams()
	const [error, setError] = React.useState(
		"[42S02][1146] (conn=107) Table 'mysql.db2' doesn't exist"
	)
	const [queryResult, setQueryResult] = React.useState('')

	const fetchLabInfo = async () => {
		try {
			const infoData = await axios.get<BasedResponse<EnrollmentInfoAPI>>(
				`/api/lab/enroll/info?enrollmentId=${params.enrollmentId}`
			)
			setEnrollmentInfo(infoData.data.data)
			setDbInfo([
				{
					label: 'Database Name',
					value: infoData.data.data.dbName ?? 'error',
				},
				{
					label: 'Database Host',
					value: infoData.data.data.dbHost ?? 'error',
				},
				{
					label: 'Database Port',
					value: infoData.data.data.dbPort ?? 'error',
				},
				{
					label: 'Database Username',
					value: infoData.data.data.dbUsername ?? 'error',
				},
				{
					label: 'Database Password',
					value: infoData.data.data.dbPassword ?? 'error',
				},
			])
		} catch {
			// alert('An error occured')
		}
	}
	const [query, setQuery] = React.useState('error')

	const handleCloseDialog = () => {
		setIsDialogOpen(false)
	}

	React.useEffect(() => {
		fetchLabInfo()
	}, [])

	const result = {
		expected_header: [
			'ID',
			'Name',
			'Age',
			'Test',
			'Test',
			'Test',
			'Test',
			'Test',
			'Test',
			'Test',
			'Test',
			'Test',
			'Test',
		],
		expected_rows: [
			['1', 'John', '20'],
			['2', 'Mary', '21'],
			['3', 'Peter', '22'],
			['3', 'Peter', '22'],
			['3', 'Peter', '22'],
			['3', 'Peter', '22'],
			['3', 'Peter', '22'],
			['3', 'Peter', '22'],
			['3', 'Peter', '22'],
			['3', 'Peter', '22'],
		],
		actual_header: ['ID', 'Name', 'Sth'],
		actual_rows: [
			['1', 'John', '20'],
			['2', 'Maryaa', '21'],
			['3', 'Peter', '22'],
		],
	}

	return (
		<Box
			sx={{
				display: 'flex',
				height: '100%',
			}}
		>
			<Box
				sx={{
					flex: 0.8,
					backgroundColor: 'white',
					boxShadow: '0px 0px 12px rgba(0, 0, 0, 0.25)',
					py: 4,
					overflowY: 'scroll',
				}}
			>
				<Box
					sx={{
						display: 'flex',
						alignItems: 'center',
						justifyContent: 'space-between',
						mb: 3.5,
						px: 4,
					}}
				>
					<Typography fontSize={24}>Spotify Test</Typography>
					<IconButton
						onClick={() => {
							setIsTask(!isTask)
						}}
					>
						<FormatListBulleted
							sx={{
								color: '#919191',
							}}
						/>
					</IconButton>
				</Box>
				{!isTask ? (
					<Box px={4}>
						{dbInfo.map((item, index) => (
							<TextField
								key={index}
								variant="outlined"
								label={item.label}
								fullWidth
								value={item.value}
								onChange={() => {}}
								sx={{
									my: 1,
								}}
							/>
						))}
					</Box>
				) : (
					<>
						{enrollmentInfo?.tasks.map((item, index) => (
							<Box key={index}>
								{index === 0 && <Divider />}
								<Box
									sx={{
										px: 4,
										py: 2.2,
										cursor: 'pointer',
										backgroundColor:
											selectedTask === index
												? '#F5F5F5'
												: 'white',
									}}
									onClick={() => {
										setSelectedTask(index)
									}}
								>
									<TaskItemList
										name={item.title}
										finished={true}
										index={index}
										selectedIndex={selectedTask}
									/>
								</Box>
								<Divider />
							</Box>
						))}
					</>
				)}
			</Box>
			<Box
				sx={{
					flex: 2,
				}}
			>
				{!isTask ? (
					<Box
						sx={{
							display: 'flex',
							justifyContent: 'center',
							alignItems: 'center',
							height: '100%',
							'& > *': {
								mx: 3,
							},
						}}
					>
						<Box
							sx={{
								width: '180px',
								height: '180px',
								borderRadius: 2,
								backgroundColor: 'white',
								boxShadow: '0px 4px 4px rgba(0, 0, 0, 0.25)',
								display: 'flex',
								flexDirection: 'column',
								justifyContent: 'center',
								alignItems: 'center',
							}}
							onClick={() => {
								setIsDialogOpen(true)
							}}
						>
							<BatchPrediction
								sx={{
									width: '70px',
									height: '90px',
									color: '#919191',
									mb: 1,
								}}
							/>
							<Typography>Mock Data</Typography>
						</Box>
						<Box
							sx={{
								width: '180px',
								height: '180px',
								borderRadius: 2,
								backgroundColor: 'white',
								boxShadow: '0px 4px 4px rgba(0, 0, 0, 0.25)',
								display: 'flex',
								flexDirection: 'column',
								justifyContent: 'center',
								alignItems: 'center',
							}}
							onClick={() => {
								!isLoading && setIsLoading(true)
							}}
						>
							<Box
								sx={{
									width: '70px',
									height: '90px',
									display: 'flex',
									justifyContent: 'center',
									alignItems: 'center',
								}}
							>
								{isLoading ? (
									<CircularProgress
										color="info"
										sx={{
											mb: 3,
										}}
									/>
								) : (
									<>
										<Autorenew
											sx={{
												width: '100%',
												height: '100%',
												color: '#919191',
											}}
										/>
									</>
								)}
							</Box>
							<Typography pt={1}>Recheck data</Typography>
						</Box>
					</Box>
				) : (
					<Box
						sx={{
							height: 'calc(100vh - 64px)',
							overflowY: 'scroll',
						}}
					>
						<Box p={4}>
							<Typography fontSize={24} fontWeight={500} mb={1.5}>
								Show first 10 rows of data from table
								'l1_tracks'
							</Typography>
							<Typography fontSize={18} mb={1.5}>
								Show first 10 rows and all columns from table
								'l1_tracks' without any condition and let MySQL
								use default sorting for the result.
							</Typography>
							<Box mb={4}>
								{taskTags.map((item, index) => (
									<Box display="flex" key={index}>
										<Typography mr={1.5}>
											{item.name}
										</Typography>
										<Typography fontWeight={700}>
											{item.value}
										</Typography>
									</Box>
								))}
							</Box>
							<TextField
								variant="outlined"
								fullWidth
								multiline
								rows={4}
								value={query}
								sx={{
									'.MuiOutlinedInput-root': {
										borderRadius: 3,
										'& fieldset': {
											borderWidth: 3,
											borderColor:
												queryResult == 'success'
													? '#4ABC4F'
													: queryResult == 'error'
													? '#BC4A4A'
													: '#a9a9a9',
										},
										'&:hover fieldset': {
											borderWidth: 3,
											borderColor:
												queryResult == 'success'
													? '#4ABC4F'
													: queryResult == 'error'
													? '#BC4A4A'
													: '#a9a9a9',
										},
									},
									'.Mui-focused fieldset': {
										borderColor:
											queryResult == 'success'
												? '#4ABC4F'
												: queryResult == 'error'
												? '#BC4A4A'
												: '#a9a9a9',
									},
									mb: 3,
								}}
								onChange={(e) => {
									setQuery(e.target.value)
								}}
							/>
							<Box
								sx={{
									backgroundColor: 'rgba(255, 118, 118, 0.2)',
									borderRadius: 3,
									p: 3,
									mb: 3,
								}}
							>
								<Typography fontFamily={'monospace'}>
									{error} {error}
								</Typography>
							</Box>
							<Box
								sx={{
									display: 'flex',
								}}
							>
								<Box
									sx={{
										flex: 1,
									}}
								>
									<Typography mb={3}>
										Expected result
									</Typography>
									<TableContainer
										sx={{
											maxWidth: 400,
										}}
									>
										<Table>
											<TableHead>
												<TableRow>
													{result.expected_header.map(
														(item, index) => (
															<TableCell
																key={index}
															>
																{item}
															</TableCell>
														)
													)}
												</TableRow>
											</TableHead>
											<TableBody>
												{result.expected_rows.map(
													(item, index) => (
														<TableRow key={index}>
															{item.map(
																(
																	item,
																	index
																) => (
																	<TableCell
																		key={
																			index
																		}
																	>
																		{item}
																	</TableCell>
																)
															)}
														</TableRow>
													)
												)}
											</TableBody>
										</Table>
									</TableContainer>
								</Box>
								<Box
									sx={{
										flex: 1,
									}}
								>
									<Typography mb={3}>
										Actual result
									</Typography>
									<TableContainer
										sx={{
											maxWidth: 400,
										}}
									>
										<Table>
											<TableHead>
												<TableRow>
													{result.actual_header.map(
														(item, index) => (
															<TableCell
																key={index}
															>
																{item}
															</TableCell>
														)
													)}
												</TableRow>
											</TableHead>
											<TableBody>
												{result.actual_rows.map(
													(item, index) => (
														<TableRow key={index}>
															{item.map(
																(
																	item,
																	index
																) => (
																	<TableCell
																		key={
																			index
																		}
																	>
																		{item}
																	</TableCell>
																)
															)}
														</TableRow>
													)
												)}
											</TableBody>
										</Table>
									</TableContainer>
								</Box>
							</Box>
						</Box>
					</Box>
				)}
			</Box>
			<Dialog open={isDialogOpen} onClose={handleCloseDialog}>
				<Box sx={style}>
					<Typography mb={2}>Generating mock data...</Typography>
					<Box
						sx={{
							height: '335px',
							backgroundColor: 'black',
							overflowY: 'scroll',
						}}
					>
						<Typography
							fontFamily={'monospace'}
							sx={{ color: 'white', p: 2 }}
						>
							<div
								dangerouslySetInnerHTML={{ __html: log }}
							></div>
						</Typography>
					</Box>
					<Button
						variant="outlined"
						sx={{
							mt: 3,
							width: '100px',
							height: '40px',
							backgroundColor: '#E1F6FB',
							color: '#0078E7',
							border: 'none',
							textTransform: 'none',
							fontSize: 16,
							':hover': {
								border: 'none',
							},
							alignSelf: 'end',
							borderRadius: 2,
						}}
						onClick={handleCloseDialog}
					>
						Finish
					</Button>
				</Box>
			</Dialog>
		</Box>
	)
}

const style = {
	position: 'fixed',
	top: '50%',
	left: '50%',
	transform: 'translate(-50%, -50%)',
	display: 'flex',
	flexDirection: 'column',
	width: 720,
	bgcolor: 'white',
	p: 4,
	borderRadius: '8px',
}

export default Lab
