# Actual outlier detection code

import logging
from concurrent.futures import ThreadPoolExecutor

import grpc
import numpy as np

from outliers_pb2 import OutliersResponse
from outliers_pb2_grpc import OutliersServicer, add_OutliersServicer_to_server

def find_outliers(data):
	"""
	return indices of data items who are greater than 2 std deviation from mean
	"""
	out = np.where(np.abs(data - data.mean()) > 2 * data.std())
	return out[0]

class OutliersServer(OutliersServicer):
	def Detect(self, request, context):
		logging.info("detect request size: {size}".format(size=len(request.metrices)))
		data = np.fromiter((m.value for m in request.metrices), dtype='float64')
		indices = find_outliers(data)
		logging.info("Found {size} outliers".format(size=len(indices)))
		resp = OutliersResponse(indices=indices)
		return resp

if __name__ == "__main__":
	logging.basicConfig(level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s")
	server = grpc.server(ThreadPoolExecutor())
	add_OutliersServicer_to_server(OutliersServer(), server)
	port = 9000
	server.add_insecure_port(f'[::]:{port}')
	server.start()
	logging.info("server ready on port {port}".format(port=port))
	server.wait_for_termination()