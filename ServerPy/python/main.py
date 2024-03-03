import os
from concurrent import futures
import grpc
from docx import Document
from pb import generator_pb2
from pb import generator_pb2_grpc

class FileGeneratorServicer(generator_pb2_grpc.FileGeneratorServicer):
    def Generate(self, request, context):
        directory = "../general/ForDownload"
        name = request.name
        if not os.path.exists(directory):
            os.makedirs(directory)
        filepath = os.path.join(directory,name+"docx")
        document = Document()
        for data in request.data:
            document.add_paragraph(data)
        document.save(filepath)
        response = generator_pb2.GenerateResponse(
            filepath=filepath,
            status="Success",
        )
        return response


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    generator_pb2_grpc.add_FileGeneratorServicer_to_server(FileGeneratorServicer(), server)
    server.add_insecure_port('[::]:9090')
    print(f"started ")
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
