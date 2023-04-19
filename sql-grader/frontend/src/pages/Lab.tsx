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
import React, { useState } from 'react'
import TaskItemList from '../components/TaskItemList'
import axios from 'axios'
import { BasedResponse } from '../types/APIs/basedResponse'
import { EnrollmentInfoAPI } from '../types/APIs/Lab/enrollment_info'
import { useParams } from 'react-router-dom'
import MockComponent from '../components/MockComponent'
import {toast} from "react-toastify";

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
	const [termKey, setKey] = useState(String(Math.random()))
	const [isTask, setIsTask] = React.useState(false)
	const [selectedTask, setSelectedTask] = React.useState(0)
	const [isLoading, setIsLoading] = React.useState(false)

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

	const clickTask = async (taskId: number) => {
		try {
			await axios.get(
				`/api/lab/task/click?enrollmentId=${params.enrollmentId}&taskId=${taskId}`
			)
		} catch {
			toast.error("Get task error")
		}
	}

	const [query, setQuery] = React.useState('error')

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
									onClick={async () => {
										setSelectedTask(index)
										await clickTask(item.id)
									}}
								>
									<TaskItemList
										name={item.title}
										finished={item.passed}
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
						<MockComponent key={termKey} />
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
									<CircularProgress color="info" />
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
		</Box>
	)
}

export default Lab
