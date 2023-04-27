import React from 'react'
import {
	CartesianGrid,
	Legend,
	Line,
	LineChart,
	ResponsiveContainer,
	Tooltip,
	XAxis,
	YAxis,
} from 'recharts'
import axios from 'axios'

const Graph = () => {
	const [data, setData] = React.useState([])

	const getData = () => {
		axios
			.get('https://wwwii.bsthun.com/mock/csworkshop/temps.json')
			.then((res) => {
				setData(mapData(res.data))
			})
	}

	const mapData = (data: any) => {
		return data.map((item: any) => {
			return {
				time: item.time.substring(11, 19),
				temp: item.value.toFixed(2),
			}
		})
	}

	React.useEffect(() => {
		getData()
	}, [])

	React.useEffect(() => {
		const interval = setInterval(() => {
			getData()
			console.log('This will run every 5 second!')
		}, 5000)
		return () => clearInterval(interval)
	}, [])

	return (
		<div
			style={{
				display: 'flex',
				width: '100vw',
				height: '100vh',
				justifyContent: 'center',
				alignItems: 'center',
			}}
		>
			<ResponsiveContainer width="80%" height="80%">
				<LineChart
					data={data}
					margin={{ top: 5, right: 30, left: 20, bottom: 40 }}
				>
					<CartesianGrid strokeDasharray="3 3" />
					<XAxis dataKey="time" angle={-45} textAnchor="end" />
					<YAxis />
					<Tooltip />
					<Legend verticalAlign="top" align="right" />
					<Line type="monotone" dataKey="temp" stroke="#8884d8" />
				</LineChart>
			</ResponsiveContainer>
		</div>
	)
}

export default Graph
