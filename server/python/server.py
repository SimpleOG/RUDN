import grpc
from concurrent import futures
import os
from grpcAPI import generator_pb2
from grpcAPI import generator_pb2_grpc
from docx import Document
class FileGeneratorServicer(generator_pb2_grpc.FileGeneratorServicer):
    def Generate(self, request, context):
        directory = "grpcAPI"
        if not os.path.exists(directory):
            os.makedirs(directory)

        filepath = os.path.join(directory, "output.docx")
        document = Document()
        for data in request.data:
            document.add_paragraph(data)
        document.save(filepath)

        response = generator_pb2.GenerateResponse()
        response.filepath = filepath
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    generator_pb2_grpc.add_FileGeneratorServicer_to_server(FileGeneratorServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()



