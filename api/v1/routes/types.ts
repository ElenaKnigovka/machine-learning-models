// types.ts

export interface ModelConfig {
    id: string;
    name: string;
    version: string;
    description?: string;
    hyperparameters: Hyperparameters;
    trainingData: TrainingData;
    createdAt: Date;
    updatedAt?: Date;
}

export interface Hyperparameters {
    learningRate: number;
    batchSize: number;
    epochs: number;
    optimizer: OptimizerType;
}

export type OptimizerType = 'adam' | 'sgd' | 'rmsprop';

export interface TrainingData {
    datasetId: string;
    features: string[];
    target: string;
    splitRatio: SplitRatio;
}

export interface SplitRatio {
    train: number;
    validation: number;
    test: number;
}

export interface ModelEvaluation {
    modelId: string;
    metrics: Metrics;
    evaluatedAt: Date;
}

export interface Metrics {
    accuracy: number;
    precision: number;
    recall: number;
    f1Score: number;
    loss: number;
}

export interface PredictionRequest {
    modelId: string;
    inputData: Record<string, any>;
}

export interface PredictionResponse {
    modelId: string;
    predictions: any[];
    timestamp: Date;
}